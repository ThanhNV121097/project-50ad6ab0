"use client";

import { useEffect, useState } from 'react';
import styles from './MessagePage.module.css';

type MessageState = 'loading' | 'ready' | 'empty' | 'error';

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? '/api';

export function MessagePage() {
  const [state, setState] = useState<MessageState>('loading');
  const [message, setMessage] = useState('');

  useEffect(() => {
    const controller = new AbortController();

    fetch(`${apiBase}/v1/message`, { signal: controller.signal })
      .then(async (response) => {
        if (response.status === 404) {
          setState('empty');
          return;
        }
        if (!response.ok) {
          setState('error');
          return;
        }
        const data = (await response.json()) as { message: string };
        setMessage(data.message);
        setState('ready');
      })
      .catch(() => setState('error'));

    return () => controller.abort();
  }, []);

  if (state === 'loading') {
    return <main className={styles.shell} aria-busy="true" aria-live="polite"><p className={styles.message}>Loading</p></main>;
  }

  if (state === 'error') {
    return <main className={styles.shell} aria-live="polite"><p className={styles.message}>Error</p></main>;
  }

  if (state === 'empty') {
    return <main className={styles.shell} aria-live="polite" />;
  }

  return <main className={styles.shell}><h1 className={styles.message}>{message}</h1></main>;
}
