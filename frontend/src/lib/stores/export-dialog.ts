import { writable } from 'svelte/store';

export const exportDialogOpen = writable(false);
export const exportData = writable<any>(null);