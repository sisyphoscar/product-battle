let products = [];
let currentWinner = null;
let currentRound = 1;
const battleResults = [];

// Main function to initialize the game
async function initializeGame() {
    const success = await fetchProducts();

    if (success) {
        renderBattleIfNeeded();
    }
}

// Fetch product data from API
async function fetchProducts() {
    const container = document.getElementById("product-container");

    try {
        const response = await fetch(BROKER_ENDPOINT + "/api/products");
        if (!response.ok) {
            throw new Error(`HTTP error: ${response.status}`);
        }

        const result = await response.json();
        if (!result || result.status !== 200 || !Array.isArray(result.data) || result.data.length < 2) {
            throw new Error("Product data format is incorrect");
        }

        products = result.data;
        currentWinner = products[0];  // default winner

        return true;
    } catch (error) {
        console.error("Get product data error:", error);
        container.innerHTML = `<p>Unable to load products, please try again later.</p>`;

        return false;
    }
}

// Render the battle UI if needed
function renderBattleIfNeeded() {
    const container = document.getElementById("product-container");
    container.innerHTML = "";  // Clear previous content

    if (products.length < 2) {
        submitBattleResults();
        renderEndMessage(container);
        return;
    }

    renderBattle(container);
}

// Submit battle results to the server
function submitBattleResults() {
    fetch(BROKER_ENDPOINT + "/api/product-battle/submit", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            game: GAME,
            roundResults: battleResults,
        }),
    });
}

// Render the end message
function renderEndMessage(container) {
    const endMessage = createElement("p", {}, "The battle has ended! Thank you for participating.");
    container.appendChild(endMessage);
}

// Render the battle UI
function renderBattle(container) {
    const challenger = products[1];
    const battleWrapper = createElement("div", { className: "battle-wrapper" });

    const winnerDiv = createProductCard(currentWinner, challenger.id, true);
    const challengerDiv = createProductCard(challenger, currentWinner.id, false);

    battleWrapper.appendChild(winnerDiv);
    battleWrapper.appendChild(challengerDiv);
    container.appendChild(battleWrapper);
}

// Create product card
function createProductCard(product, opponentId, isCurrentWinner) {
    const productDiv = createElement("div");
    const title = createElement("h2", {}, product.name);
    const image = createElement("img", { src: product.imageUrl, alt: product.name, className: "product-image" });

    // if image fails to load, set a default image
    image.onerror = () => {
        image.src = "/static/images/no-image.png";
        image.alt = "No image available";
    };

    const description = createElement("p", {}, product.description);
    const button = createElement("button", {}, "Vote");

    button.onclick = () => vote(product.id, opponentId, isCurrentWinner);

    productDiv.appendChild(image);
    productDiv.appendChild(title);
    productDiv.appendChild(description);
    productDiv.appendChild(button);

    return productDiv;
}

// Handle voting
function vote(winnerId, loserId, isCurrentWinner) {
    battleResults.push({
        round: currentRound,
        winnerId: winnerId,
        loserId: loserId,
    });

    currentRound++;

    if (!isCurrentWinner) {
        currentWinner = products[1];
    }

    // Remove the loser from the products
    products = products.filter(product => product.id !== loserId);

    renderBattleIfNeeded();
}

// Create an HTML element with attributes and text content
function createElement(tag, attributes = {}, textContent = "") {
    const element = document.createElement(tag);
    for (const [key, value] of Object.entries(attributes)) {
        element[key] = value;
    }
    if (textContent) {
        element.textContent = textContent;
    }
    return element;
}

// Call the main function to start the game
initializeGame();