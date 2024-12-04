<script lang="ts">
    import type { Selected } from "@skeletonlabs/skeleton";
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
        </CardHeader>
        <CardContent class="space-y-4">
            <div class="space-y-2">
                <Select bind:selected={selectedFormat}>
                    <SelectTrigger class="z-[10001]">
                        <SelectValue placeholder="Select a format" />
                    </SelectTrigger>
                    <SelectContent class="z-[10001]">
                        <SelectItem value="geojson">GeoJSON</SelectItem>
                        <SelectItem value="kml">KML</SelectItem>
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
                        Split geometry
                    </label>
                    <Switch id="split-geometry" bind:checked={splitGeometry} />
                </div>
            </div>
            <Button class="w-full" on:click={downloadGeoJSON}>
                Download {selectedFormat?.value ?? ''}
            </Button>
        </CardContent>
        <CardFooter>
            <Button variant="outline" class="w-full" on:click={() => $exportDialogOpen = false}>
                Cancel
            </Button>
        </CardFooter>
    </Card>
</div>
{/if}