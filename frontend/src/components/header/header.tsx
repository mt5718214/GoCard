import { component$, useStylesScoped$ } from '@builder.io/qwik';
import styles from './header.css?inline';

export default component$(() => {
  useStylesScoped$(styles);

  return (
    <header>
      <div class="logotext">
        <a href="/">
          <h1 class="px-2 text-dark">TodoList</h1>
        </a>
      </div>
    </header>
  );
});
