<script>
	import { AppShell } from '@skeletonlabs/skeleton';
	import Sidebar from './components/Sidebar.svelte'
	import ChatHeader from './components/ChatHeader.svelte'
	import ChatFooter from './components/ChatFooter.svelte'
	import { selectedPeer } from '$lib/stores.ts'
	
</script>

<AppShell>
	<svelte:fragment slot="sidebarLeft">
		<Sidebar />
	</svelte:fragment>

		<svelte:fragment slot="pageHeader">
			{#if $selectedPeer}
				<ChatHeader />
			{/if}
		</svelte:fragment>
		
		<slot />
	

	<svelte:fragment slot="pageFooter">
		{#if $selectedPeer}
			<ChatFooter />
		{/if}
	</svelte:fragment>
</AppShell>