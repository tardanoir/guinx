<script lang="ts">
    import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "$lib/components/ui/card";
    import { Button } from "$lib/components/ui/button";
    import { importDialogOpen } from '$lib/stores/import-dialog';
    import { kml } from '@tmcw/togeojson';

    let fileInput: HTMLInputElement;
    let geometryText = '';

    async function handleFileSelect(event: Event) {
        const target = event.target as HTMLInputElement;
        const file = target?.files?.[0];
        
        if (!file) return;

        try {
            let geoJson;
            
            if (file.name.toLowerCase().endsWith('.kml')) {
                // Handle KML
                const text = await file.text();
                const parser = new DOMParser();
                const kmlDoc = parser.parseFromString(text, 'text/xml');
                geoJson = kml(kmlDoc);
            } else if (file.name.toLowerCase().endsWith('.geojson')) {
                // Handle GeoJSON
                const text = await file.text();
                geoJson = JSON.parse(text);
                
                // Ensure it's a FeatureCollection if it's not already
                if (geoJson.type !== 'FeatureCollection') {
                    geoJson = {
                        type: 'FeatureCollection',
                        features: geoJson.type === 'Feature' ? [geoJson] : []
                    };
                }
            } else {
                throw new Error('Unsupported file format');
            }
            
            // Dispatch the event with the processed GeoJSON
            window.dispatchEvent(new CustomEvent('import-geometry', {
                detail: { data: geoJson }
            }));
            
            $importDialogOpen = false;
        } catch (error) {
            console.error('Error importing file:', error);
            alert('Error importing file. Please make sure it\'s a valid GeoJSON or KML file.');
        }
    }

    function handleTextImport() {
        if (!geometryText.trim()) {
            alert('Please paste some geometry data first');
            return;
        }

        try {
            let geoJson;
            
            // Try parsing as JSON first
            try {
                geoJson = JSON.parse(geometryText);
                
                // Ensure it's a FeatureCollection
                if (geoJson.type !== 'FeatureCollection') {
                    geoJson = {
                        type: 'FeatureCollection',
                        features: geoJson.type === 'Feature' ? [geoJson] : []
                    };
                }
            } catch {
                // If JSON parsing fails, try KML
                const parser = new DOMParser();
                const kmlDoc = parser.parseFromString(geometryText, 'text/xml');
                geoJson = kml(kmlDoc);
            }

            // Dispatch the event with the processed GeoJSON
            window.dispatchEvent(new CustomEvent('import-geometry', {
                detail: { data: geoJson }
            }));
            
            // Clear the textarea and close the dialog
            geometryText = '';
            $importDialogOpen = false;
        } catch (error) {
            console.error('Error importing text:', error);
            alert('Error importing text. Please make sure it\'s valid GeoJSON or KML.');
        }
    }
</script>

{#if $importDialogOpen}
<div class="fixed inset-0 bg-black/50 flex items-center justify-center z-[9999]">
    <Card class="w-[350px] relative z-[10000]">
        <CardHeader>
            <CardTitle>Import Geometry</CardTitle>
            <CardDescription>Import GeoJSON or KML files</CardDescription>
            <button
                class="absolute top-2 right-2 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
                on:click={() => $importDialogOpen = false}
            >
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                    <path d="M18 6 6 18"/>
                    <path d="m6 6 12 12"/>
                </svg>
                <span class="sr-only">Close</span>
            </button>
        </CardHeader>
        <CardContent>
            <div class="space-y-4">
                <button 
                    type="button"
                    class="w-full border-2 border-dashed border-muted-foreground/25 rounded-lg p-8 text-center hover:bg-accent/50 transition-colors"
                    on:dragover|preventDefault
                    on:drop|preventDefault={handleFileSelect}
                    on:click={() => fileInput.click()}
                >
                    <div class="text-muted-foreground mb-2">
                        Drag and drop your file here or click to browse
                    </div>
                    <div class="text-xs text-muted-foreground/75">
                        Supported formats: .geojson, .kml
                    </div>
                    <input
                        type="file"
                        accept=".geojson,.kml"
                        bind:this={fileInput}
                        on:change={handleFileSelect}
                        class="hidden"
                    />
                </button>

                <div class="space-y-2">
                    <label for="geometry-text" class="text-sm text-muted-foreground">
                        Or paste your GeoJSON/KML here:
                    </label>
                    <textarea
                        id="geometry-text"
                        class="w-full min-h-[100px] rounded-md border bg-transparent px-3 py-2 text-sm shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring mb-2"
                        placeholder="Paste your geometry data here..."
                        bind:value={geometryText}
                    ></textarea>
                    <Button 
                        class="w-full"
                        on:click={handleTextImport}
                    >
                        Import Text
                    </Button>
                </div>
            </div>
        </CardContent>
    </Card>
</div>
{/if}