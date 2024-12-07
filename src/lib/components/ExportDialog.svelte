<script lang="ts">
    import type { Selected } from "bits-ui";
    import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "$lib/components/ui/card";
    import { Button } from "$lib/components/ui/button";
    import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "$lib/components/ui/select";
    import { exportDialogOpen, exportData } from '$lib/stores/export-dialog';
    import Switch from "$lib/components/ui/switch/switch.svelte";

    let selectedFormat: Selected<string>;
    let splitGeometry = false;

    function downloadGeoJSON() {
        if (!$exportData) return;
        
        if (splitGeometry) {
            // Assuming exportData is a FeatureCollection, split by features
            $exportData.features.forEach((feature: any, index: number) => {
                const singleFeatureCollection = {
                    type: "FeatureCollection",
                    features: [feature]
                };
                
                const dataStr = JSON.stringify(singleFeatureCollection);
                const dataUri = 'data:application/json;charset=utf-8,' + encodeURIComponent(dataStr);
                const linkElement = document.createElement('a');
                linkElement.setAttribute('href', dataUri);
                linkElement.setAttribute('download', `map-data-${index + 1}.geojson`);
                linkElement.click();
            });
        } else {
            // Original single file download logic
            const dataStr = JSON.stringify($exportData);
            const dataUri = 'data:application/json;charset=utf-8,' + encodeURIComponent(dataStr);
            const linkElement = document.createElement('a');
            linkElement.setAttribute('href', dataUri);
            linkElement.setAttribute('download', 'map-data.geojson');
            linkElement.click();
        }
        
        $exportDialogOpen = false;
    }
</script>

{#if $exportDialogOpen}
<div class="fixed inset-0 bg-black/50 flex items-center justify-center z-[9999]">
    <Card class="w-[350px] relative z-[10000]">
        <CardHeader>
            <CardTitle>Export Geometry</CardTitle>
            <CardDescription>Download your geometry data</CardDescription>
            <button
                class="absolute top-2 right-2 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
                on:click={() => $exportDialogOpen = false}
            >
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                    <path d="M18 6 6 18"/>
                    <path d="m6 6 12 12"/>
                </svg>
                <span class="sr-only">Close</span>
            </button>
        </CardHeader>
        <CardContent class="space-y-4">
            <div class="space-y-2">
                <Select bind:selected={selectedFormat}>
                    <SelectTrigger class="z-[10001]">
                        <SelectValue placeholder="Select a format" />
                    </SelectTrigger>
                    <SelectContent class="z-[10001]">
                        <SelectItem value=".geojson">GeoJSON</SelectItem>
                        <SelectItem value=".kml">KML</SelectItem>
                    </SelectContent>
                </Select>
                <div class="flex items-center justify-between">
                    <label for="include-metadata" class="text-sm font-medium leading-none">
                        Include metadata
                    </label>
                    <Switch id="include-metadata" />
                </div>
                <div class="flex items-center justify-between">
                    <label for="split-geometry" class="text-sm font-medium leading-none">
                        Download single geometries
                    </label>
                    <Switch id="split-geometry" bind:checked={splitGeometry} />
                </div>
            </div>
            <div class="flex flex-row gap-2">
                <Button 
                    variant="outline"
                    class="w-full"
                    on:click={() => $exportDialogOpen = false}
                >
                    Cancel
                </Button>
                <Button 
                    class="w-full"
                    on:click={downloadGeoJSON}
                >
                    Download Geometries
                </Button>
        </CardContent>
    </Card>
</div>
{/if}