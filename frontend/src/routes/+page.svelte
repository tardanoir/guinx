<script lang="ts">
    import Map from '$lib/components/Map.svelte';
</script>
  
<div class="layout">
    <link href='https://api.mapbox.com/mapbox-gl-js/v2.14.1/mapbox-gl.css' rel='stylesheet' />
    <Map />
</div>
  
<style>
    .layout {
        height: 100vh;
        width: 100vw;
        position: relative;
    }
</style>