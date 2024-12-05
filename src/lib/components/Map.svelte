<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import { browser } from '$app/environment';
    import ExportDialog from '$lib/components/ExportDialog.svelte';
    import { exportDialogOpen, exportData } from '$lib/stores/export-dialog';
    import { importDialogOpen } from '$lib/stores/import-dialog';
    import ImportDialog from '$lib/components/ImportDialog.svelte';
    import mapboxgl from 'mapbox-gl';
    import MapboxDraw from '@mapbox/mapbox-gl-draw';
    import MapControls from './MapControls.svelte';

    // You'll need to get a Mapbox access token
    mapboxgl.accessToken = 'pk.eyJ1IjoidGFyZGFub2lyIiwiYSI6ImNtNGJmdXFxeTAzNzMycW5kZHpyNWo3Y2kifQ.tBnUr1i-Pgcmbki6mCgRzA';
    let mapElement: HTMLElement;
    let map: mapboxgl.Map;
    let draw: any;
    let status = 'saved';
    let mode = 'Normal';
    let currentBranch = 'main';

    onMount(async () => {
        if (browser) {
            map = new mapboxgl.Map({
                container: mapElement,
                style: 'mapbox://styles/mapbox/satellite-v9',
                center: [-51.697926108060145, -31.045980875783325],
                zoom: 8
            });

            draw = new MapboxDraw({
                displayControlsDefault: false,
                controls: {
                    polygon: false,
                    line_string: false,
                    point: false,
                    trash: false 
                }
            });

            map.addControl(draw);

            map.on('draw.create', updateDrawings);
            map.on('draw.delete', updateDrawings);
            map.on('draw.update', updateDrawings);

            window.addEventListener('import-geometry', handleImportGeometry as EventListener);
        }
    });

    function updateDrawings() {
        // This function will be called whenever drawings are modified
        const data = draw.getAll();
        // You can handle the updated data here if needed
    }

    function handleImportGeometry(e: CustomEvent) {
        const geoJson = e.detail.data;
        draw.deleteAll();
        draw.add(geoJson);
    }

    function handleExport() {
        if (draw) {
            $exportData = draw.getAll();
            $exportDialogOpen = true;
        }
    }

    function handleImport() {
        $importDialogOpen = true;
    }

    function zoomIn() {
        map.zoomIn();
    }

    function zoomOut() {
        map.zoomOut();
    }

    function drawLineString() {
        draw.changeMode('draw_line_string');
    }

    function drawPolygon() {
        draw.changeMode('draw_polygon');
    }

    function drawPoint() {
        draw.changeMode('draw_point');
    }

    function drawTrash() {
        draw.deleteAll();
    }

    onDestroy(() => {
        if (browser) {
            window.removeEventListener('import-geometry', handleImportGeometry as EventListener);
            if (map) {
                map.remove();
            }
        }
    });
</script>

<div class="map-container">
    <div bind:this={mapElement} id="map"></div>
    <MapControls {draw} {map} />
    <div class="bottom-bar">
        <div class="status-item">
            {#if status === 'saved'}
                <span class="status-icon">✓</span>
                <span>Saved</span>
            {:else}
                <span class="status-icon">●</span>
                <span>Unsaved changes</span>
            {/if}
        </div>
        <div class="status-item">
            <span>{mode}</span>
        </div>
        <div class="status-item">
            <span class="branch-icon">⎇</span>
            <span>{currentBranch}</span>
        </div>
    </div>
    <div class="bottom-status"></div>
</div>
<ExportDialog />
<ImportDialog />


<style>
    .map-container {
        position: relative;
        height: 100vh;
        width: 100vw;
        overflow: hidden;
        display: flex;
        flex-direction: column;
    }
  
    #map {
        flex: 1;
        width: 100%;
        min-height: 0;
    }

    .bottom-bar {
        position: absolute;
        bottom: 0;
        width: 100%;
        height: 34px;
        background-color: black;
        z-index: 1000;
    }

    .bottom-status {
        position: absolute;
        bottom: 0;
        width: 100%;
        height: 4px;
        background-color: #2563EB;
        z-index: 1000;
    }

    .bottom-bar .status-item {
        display: inline-block;
        padding: 0 5px;
        position: relative;
    }


    /* Adjust styles for Mapbox attribution */
    :global(.mapboxgl-ctrl-bottom-left),
    :global(.mapboxgl-ctrl-bottom-right) {
        position: absolute;
        bottom: 38px;
        z-index: 1001;
    }

    /* Ensure the map container fills the viewport */
    .map-container {
        position: relative;
        height: 100vh;
        width: 100vw;
        overflow: hidden;
        display: flex;
        flex-direction: column;
    }

    #map {
        flex: 1;
        width: 100%;
        min-height: 0;
    }

    :global([data-theme='dark'] .mapboxgl-ctrl-icon) {
        --dark-mode-invert: 1;
    }
</style>