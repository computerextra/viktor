<script>
	import Geburtstagsliste from '$lib/Geburtstagsliste.svelte';
</script>

<div class="container mx-auto mt-12">
	<div class="grid grid-cols-1 gap-12 lg:grid-cols-2">
		<div>
			<div class="space-y-8">
				<h1 class="text-6xl font-extrabold leading-none tracking-tighter text-black md:text-8xl">
					Tach auf
					<br />
					<span class="inline-block -rotate-3 transform bg-black px-4 py-2 text-white">Viktor</span>
				</h1>
				<p class="max-w-lg font-mono text-xl text-gray-200">
					Das ist der neue Viktor. Einige Funktionen sind experimentell, bei Problemen bitte mir
					Bescheid geben.
				</p>
				<a
					href="mailto:johannes.kirchner@computer-extra.de?subject=Ich habe ein Problem mit Viktor&body=Mein Problem: %0D%0DWas habe ich gemacht, dass das Problem auftritt? : %0D%0DWann tritt das Problem auf? : %0D%0DWie oft ist das Problem bereits aufgetreten? : %0D%0DEntsteht das Problem jedes Mal? :%0D%0D"
					id="bescheid"
					class="focus:outline-hidden inline-flex items-center gap-x-2 rounded-lg border border-transparent bg-blue-600 px-4 py-3 text-sm font-medium text-white hover:bg-blue-700 focus:bg-blue-700 disabled:opacity-50"
				>
					Bescheid!
				</a>
			</div>
		</div>
	</div>
	<Geburtstagsliste />
</div>
