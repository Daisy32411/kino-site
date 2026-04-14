let lastScroll = 0;
const header = document.querySelector("header");
const scrollThreshold = 10;

window.addEventListener("scroll", () => {
    let currentScroll = window.pageYOffset;

    if (Math.abs(currentScroll - lastScroll) < scrollThreshold) return;

    if (currentScroll > lastScroll && currentScroll > 100) {
        header.classList.add("hide");
    } else {
        header.classList.remove("hide");
    }

    lastScroll = currentScroll;
});