import { component$ } from '@builder.io/qwik';
import type { DocumentHead } from '@builder.io/qwik-city';
import { Link } from '@builder.io/qwik-city';

export default component$(() => {
  return (
    <div>
      <Link href='/login'>login</Link>
    </div>
  );
});

export const head: DocumentHead = {
  title: 'todo-list',
};
