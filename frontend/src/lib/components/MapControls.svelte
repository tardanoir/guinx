<script lang="ts">
    import Button from "$lib/components/ui/button/button.svelte";
    import { importDialogOpen } from '$lib/stores/import-dialog';
    import { exportDialogOpen, exportData } from '$lib/stores/export-dialog';
    export let draw: any;
    export let map: mapboxgl.Map;

    function handleImport() {
        $importDialogOpen = true;
    }

    function handleExport() {
        if (draw) {
            $exportData = draw.getAll();
            $exportDialogOpen = true;
        }
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
</script>

<div class="import-export-buttons">
    <Button 
        variant="outline"
        class="bg-background hover:bg-accent text-foreground shadow-sm"
        on:click={handleImport}
    >
        <img src="/icons/upload.svg" alt="Upload" />
    </Button>
    <Button 
        variant="outline"
        class="bg-background hover:bg-accent text-foreground shadow-sm"
        on:click={handleExport}
    >
        <img src="/icons/download.svg" alt="Download" />
    </Button>
</div>
<div class="zoom-controls">
    <Button 
        variant="outline"
        class="bg-background hover:bg-accent text-foreground shadow-sm"
        on:click={zoomIn}
    >
        <img src="/icons/zoomIn.svg" alt="Zoom In" />
    </Button>
    <Button 
        variant="outline"
        class="bg-background hover:bg-accent text-foreground shadow-sm"
        on:click={zoomOut}
    >
        <img src="/icons/zoomOut.svg" alt="Zoom Out" />
    </Button>
</div>
<div class="drawing-controls">
    <Button 
        variant="outline"
        class="bg-background hover:bg-accent text-foreground shadow-sm"
        on:click={drawLineString}
    >
        <img src="/icons/linestring.svg" alt="Draw Line String" />
    </Button>
    <Button 
        variant="outline"
        class="bg-background hover:bg-accent text-foreground shadow-sm"
        on:click={drawPolygon}
    >
        <img src="/icons/polygon.svg" alt="Draw Polygon" />
    </Button>
    <Button 
        variant="outline"
        class="bg-background hover:bg-accent text-foreground shadow-sm"
        on:click={drawPoint}
    >
        <img src="/icons/point.svg" alt="Draw Point" />
    </Button>
    <Button 
        variant="outline"
        class="bg-background hover:bg-accent text-foreground shadow-sm"
        on:click={drawTrash}
    >
        <img src="/icons/trash.svg" alt="Trash" />
    </Button>
</div>

<style>
    .import-export-buttons {
        position: absolute;
        top: 300px;
        right: 10px;
        display: flex;
        flex-direction: column;
        gap: 3px;
    }

    .zoom-controls {
        position: absolute;
        top: 10px;
        right: 10px;
        display: flex;
        flex-direction: column;
        gap: 3px;
    }

    .drawing-controls {
        position: absolute;
        top: 120px;
        right: 10px;
        display: flex;
        flex-direction: column;
        gap: 3px;
    }
</style>