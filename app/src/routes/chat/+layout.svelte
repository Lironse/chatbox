<script>
	import { AppShell } from '@skeletonlabs/skeleton';
	import Sidebar from './components/Sidebar.svelte'
	import ChatHeader from './components/ChatHeader.svelte'
	import ChatFooter from './components/ChatFooter.svelte'
	
</script>

<AppShell>
	<svelte:fragment slot="sidebarLeft">
		<Sidebar />
	</svelte:fragment>

	<svelte:fragment slot="pageHeader">
		<ChatHeader />
	</svelte:fragment>

	<slot />

	<svelte:fragment slot="pageFooter">
		<ChatFooter />
	</svelte:fragment>
</AppShell>