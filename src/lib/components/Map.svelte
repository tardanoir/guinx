<script lang="ts">
    import { onMount } from 'svelte';
    import { browser } from '$app/environment';
    import Button from "$lib/components/ui/button/button.svelte";
    
    let mapElement: HTMLElement;
    let map: any;
    let drawnItems: any;
    let L: any;

    onMount(async () => {
        if (browser) {
            const leaflet = await import('leaflet');
            L = leaflet.default;
            await import('leaflet-draw');
            
            map = L.map(mapElement).setView([-27.5954, -48.5480], 13);
            
            L.tileLayer('https://{s}.google.com/vt/lyrs=y&x={x}&y={y}&z={z}', {
                subdomains: ['mt0', 'mt1', 'mt2', 'mt3'],
                attribution: '© OpenStreetMap contributors'
            }).addTo(map);
            
            drawnItems = new L.FeatureGroup();
            map.addLayer(drawnItems);
            
            const drawControl = new (L.Control as any).Draw({
                edit: {
                    featureGroup: drawnItems
                },
                draw: {
                    polygon: true,
                    polyline: true,
                    rectangle: true,
                    circle: false,
                    circlemarker: false,
                    marker: true
                }
            });
            map.addControl(drawControl);
            
            map.on('draw:created', (e: any) => {
                const layer = e.layer;
                drawnItems.addLayer(layer);
            });
        }
    });

    function exportGeoJSON() {
        if (drawnItems) {
            const data = drawnItems.toGeoJSON();
            const dataStr = JSON.stringify(data);
            const dataUri = 'data:application/json;charset=utf-8,'+ encodeURIComponent(dataStr);
            
            const exportFileDefaultName = 'map-data.geojson';
            const linkElement = document.createElement('a');
            linkElement.setAttribute('href', dataUri);
            linkElement.setAttribute('download', exportFileDefaultName);
            linkElement.click();
        }
    }
</script>

<div class="map-container">
    <div bind:this={mapElement} id="map"></div>
    <div class="export-button-container">
        <Button 
            variant="outline"
            class="bg-background hover:bg-accent text-foreground shadow-sm"
            on:click={exportGeoJSON}
        >
            Export GeoJSON
        </Button>
    </div>
</div>

<style>
    .map-container {
        position: relative;
        height: 100%;
        width: 100%;
    }
  
    #map {
        height: 100%;
        width: 100%;
    }

    .export-button-container {
        position: absolute;
        top: 20px;
        right: 20px;
        z-index: 9999;
    }

    :global(.leaflet-container) {
        height: 100%;
        width: 100%;
    }
</style>