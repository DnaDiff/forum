class Category extends HTMLElement {
  constructor(ID, name, description, posts) {
    super();

    this.ID = ID;
    this.name = name;
    this.description = description;
    this.posts = posts;

    // Create HTML elements and set values
    this.dataset.id = ID;
    this.classList.add("category");
    this.innerHTML = `
    <header class="category__header">
      <div class="category__name">${this.name}</div>
      <div class="category__description">${this.description}</div>
    </header>
    <section class="category__body">
      <div class="category__posts"></div>
    </section>
    <footer class="category__footer">
      <button class="category__button">Show more</button>
    </footer>
    `;

    // Add posts to the category
    this.posts.forEach((post) => {
      const postElement = new Post(
        post.ID,
        post.parentID,
        post.title,
        post.content,
        post.date,
        post.comments,
        post.rating,
        post.userID,
        post.username,
        post.userAvatar
      );
      this.querySelector(".category__posts").appendChild(postElement);
    });
  }
}
