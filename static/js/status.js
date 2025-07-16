async function getStatus() {
    const url = "https://status.computer-extra.net/status.php"
    const res = await fetch(url)
    const json = await res.json()
    const data = json[0]
    // const since = data.since
    const status = data.status

    // let active = ""
    let statRes = ""
    switch(status) {
        case "Anwesend":
            statRes = `<span class="flex flex-row items-center gap-x-2">
             Johannes Status: Anwesend <svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-check-check-icon lucide-check-check !text-green-500"
		><path d="M18 6 7 17l-5-5" class="!text-green-500"></path><path class="!text-green-500" d="m22 10-7.5 7.5L13 16"></path></svg>
            </span>`
            break
        case "Abwesend":
             statRes = `<span class="flex flex-row items-center gap-x-2">
             Johannes Status: Abwesend <svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-ban-icon lucide-ban  !text-red-500"
		>
			<circle class="!text-red-500" cx="12" cy="12" r="10"></circle>
			<path class="!text-red-500" d="m4.9 4.9 14.2 14.2"></path>
		</svg></span>`
            break
        case "Beschäftigt":
             statRes = `<span class="flex flex-row items-center gap-x-2">
             Johannes Status: Beschäftigt <svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-bell-off-icon lucide-bell-off  !text-red-500"
		>
			<path class="!text-red-500" d="M10.268 21a2 2 0 0 0 3.464 0"></path>
			<path class="!text-red-500" d="M17 17H4a1 1 0 0 1-.74-1.673C4.59 13.956 6 12.499 6 8a6 6 0 0 1 .258-1.742"></path>
			<path class="!text-red-500" d="m2 2 20 20"></path>
			<path class="!text-red-500" d="M8.668 3.01A6 6 0 0 1 18 8c0 2.687.77 4.653 1.707 6.05"></path>
		</svg></span>`
            break
        case "Am Telefonieren":
             statRes = `<span class="flex flex-row items-center gap-x-2">
             Johannes Status: Am Telefonieren <svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-phone-call-icon lucide-phone-call  !text-blue-500"
		>
			<path class="!text-blue-500" d="M13 2a9 9 0 0 1 9 9"></path>
			<path class="!text-blue-500" d="M13 6a5 5 0 0 1 5 5"></path>
			<path
				class="!text-blue-500"
				d="M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"
			></path>
		</svg></span>`
            break
        case "Im Mittag":
             statRes = `<span class="flex flex-row items-center gap-x-2">
             Johannes Status: Im Mittag <svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-utensils-crossed-icon lucide-utensils-crossed  !text-red-500"
		>
			<path class="!text-red-500" d="m16 2-2.3 2.3a3 3 0 0 0 0 4.2l1.8 1.8a3 3 0 0 0 4.2 0L22 8"></path>
			<path class="!text-red-500" d="M15 15 3.3 3.3a4.2 4.2 0 0 0 0 6l7.3 7.3c.7.7 2 .7 2.8 0L15 15Zm0 0 7 7"></path>
			<path class="!text-red-500" d="m2.1 21.8 6.4-6.3"></path>
			<path class="!text-red-500" d="m19 5-7 7"></path>
		</svg></span>`
            break
        case "Am Rauchen":
             statRes = `<span class="flex flex-row items-center gap-x-2">
             Johannes Status: Am Rauchen <svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-cigarette-icon lucide-cigarette  !text-red-500"
		>
			<path class="!text-red-500" d="M17 12H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h14"></path>
			<path class="!text-red-500" d="M18 8c0-2.5-2-2.5-2-5"></path>
			<path class="!text-red-500" d="M21 16a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1"></path>
			<path class="!text-red-500" d="M22 8c0-2.5-2-2.5-2-5"></path>
			<path class="!text-red-500" d="M7 12v4"></path>
		</svg></span>`
            break
        case "Auf Silbers Platz":
             statRes = `<span class="flex flex-row items-center gap-x-2">
             Johannes Status: Auf Silbers Platz <svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-utensils-crossed-icon lucide-utensils-crossed  !text-red-500"
		>
			<path class="!text-red-500" d="m16 2-2.3 2.3a3 3 0 0 0 0 4.2l1.8 1.8a3 3 0 0 0 4.2 0L22 8"></path>
			<path class="!text-red-500" d="M15 15 3.3 3.3a4.2 4.2 0 0 0 0 6l7.3 7.3c.7.7 2 .7 2.8 0L15 15Zm0 0 7 7"></path>
			<path class="!text-red-500" d="m2.1 21.8 6.4-6.3"></path>
			<path class="!text-red-500" d="m19 5-7 7"></path>
		</svg></span>`
            break
        case "In Wartung":
             statRes = `<span class="flex flex-row items-center gap-x-2">
             Johannes Status: In Wartung <svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-server-icon lucide-server  !text-red-500"
		>
			<rect class="!text-red-500" width="20" height="8" x="2" y="2" rx="2" ry="2"></rect>
			<rect class="!text-red-500" width="20" height="8" x="2" y="14" rx="2" ry="2"></rect>
			<line class="!text-red-500" x1="6" x2="6.01" y1="6" y2="6"></line>
			<line class="!text-red-500" x1="6" x2="6.01" y1="18" y2="18"></line>
		</svg></span>`
            break
    }

    // if(active.length > 0) {
    //     const elems = document.querySelectorAll("[data-status]")
        
    //     elems.forEach(x => {
    //         if(x.dataset.status == active) {
    //             x.classList.remove("hidden")
    //         }else {
    //             x.classList.add("hidden")
    //         }
    //     })
    // }

    const span = document.querySelector("#status")
    span.innerHTML = statRes
    // const sinceSpan = document.querySelector("#since")
    // sinceSpan.innerHTML = `Seit: ${since} Uhr`
}

getStatus()

setInterval(async () => {
    await getStatus()
}, 5 * 1000)