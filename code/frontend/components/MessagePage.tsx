"use client";

import type { MessageResponse } from '../lib/mock/store-and-serve-message';
import styles from './MessagePage.module.css';

export function MessagePage({ response }: { response: MessageResponse }) {
  if (response.state === 'loading') {
    return (
      <main className={styles.shell} aria-busy="true" aria-live="polite">
        <p className={styles.message}>Loading</p>
      </main>
    );
  }

  if (response.state === 'error') {
    return (
      <main className={styles.shell} aria-live="polite">
        <p className={styles.message}>Error</p>
      </main>
    );
  }

  if (response.state === 'empty') {
    return (
      <main className={styles.shell} aria-live="polite">
        <p className={styles.message}>No message</p>
      </main>
    );
  }

  return (
    <main className={styles.shell}>
      <h1 className={styles.message}>{response.message}</h1>
    </main>
  );
}
