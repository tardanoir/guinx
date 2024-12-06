<script lang="ts">
    import 'mapbox-gl/dist/mapbox-gl.css';
    import '@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css';
    import "../app.pcss";
    import { isAuthenticated } from '$lib/stores/auth';

    if (typeof document !== 'undefined') {
        document.documentElement.classList.add('dark');
    }
</script>

<slot />