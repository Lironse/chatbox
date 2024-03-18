<script lang="ts">
	import SettingsCard from './SettingsCard.svelte';
	import ProfileCard from './ProfileCard.svelte';
	import AddPeerButton from './AddPeerButton.svelte'
	import PeerList from './PeerList.svelte';
	import SettingsButton from './SettingsButton.svelte';

</script>

<div class="w-full h-full flex flex-col p-4 gap-5 border-r border-gray-500">
	<ProfileCard />

	<PeerList />

	<AddPeerButton />
	
	<SettingsButton />

	<SettingsCard />
</div>
