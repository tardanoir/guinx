<script>
    import 'mapbox-gl/dist/mapbox-gl.css';
    import '@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css';
    import "../app.pcss";

    if (typeof document !== 'undefined') {
        document.documentElement.classList.add('dark');
    }
</script>

<slot />
