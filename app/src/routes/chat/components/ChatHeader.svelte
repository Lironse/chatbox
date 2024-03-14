<script lang='ts'>
	import type { Peer } from '$lib/peer';
    import { selectedPeer } from '$lib/stores.ts'
    import { Avatar } from '@skeletonlabs/skeleton'

    let peer: Peer
    selectedPeer.subscribe((value) => {
        peer = value
    })

</script>

<div class='flex flex-row gap-5 justify-start items-center'>
    <Avatar src="https://i.pravatar.cc/?img={peer?.avatar}" width="w-12" />
    <span class='h2 font-bold align-middle'>{peer?.name}</span>
</div>