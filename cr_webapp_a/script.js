document.addEventListener("DOMContentLoaded", function () {
  const numberOfCards = 8; // Hard code the number of cards you want to add
  const container = document.querySelector(".container");
  const componentsPerRow = 4; // Number of cards per row
  let row; // Variable to hold the current row

  for (let i = 0; i < numberOfCards; i++) {
    // Create a new card element
    const card = document.createElement("div");
    card.className = "one-fourth column card"; // Apply Skeleton CSS classes

    // Set the HTML content for the card
    card.innerHTML = `
            <div class="pokemon-ditto"></div>
            <h3 class="card-title">Card Title ${i + 1}</h3>
            <p class="card-description">This is card number ${
              i + 1
            }. Customize it as needed.</p>
        `;

    // If this is the first card in the row, or if the current row is full, create a new row
    if (i % componentsPerRow === 0) {
      row = document.createElement("div");
      row.className = "row"; // Apply Skeleton CSS class for row
      container.appendChild(row);
    }

    // Append the card to the current row
    row.appendChild(card);
  }
});
