<script lang="ts">
    let status = 'saved';
    let mode = 'Insert';
    let currentBranch = 'main';
</script>

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

<style>
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

</style>