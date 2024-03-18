<script>
	import { AppShell } from '@skeletonlabs/skeleton';
	import Sidebar from './components/Sidebar.svelte';
</script>

<AppShell>
	<svelte:fragment slot="sidebarLeft">
		<Sidebar />
	</svelte:fragment>

	<slot />
</AppShell>
