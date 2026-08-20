"use client";

import { useEffect, useState } from 'react';
import styles from './MessagePage.module.css';

type MessageResponse =
  | { state: 'loading' }
  | { state: 'empty' }
  | { state: 'error' }
  | { state: 'ready'; message: string };

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? '/api';

export function MessagePage() {
  const [response, setResponse] = useState<MessageResponse>({ state: 'loading' });

  useEffect(() => {
    const controller = new AbortController();

    async function loadMessage() {
      try {
        const res = await fetch(`${apiBase}/v1/message`, {
          signal: controller.signal,
          headers: { Accept: 'application/json' },
        });
        const data = await res.json().catch(() => null);

        if (!res.ok) {
          setResponse(res.status === 404 ? { state: 'empty' } : { state: 'error' });
          return;
        }

        const message = typeof data?.message === 'string' ? data.message.trim() : '';
        setResponse(message ? { state: 'ready', message } : { state: 'empty' });
      } catch {
        setResponse({ state: 'error' });
      }
    }

    loadMessage();
    return () => controller.abort();
  }, []);

  return (
    <main className={styles.shell} aria-busy={response.state === 'loading'} aria-live="polite">
      <h1 className={styles.message}>{response.state === 'ready' ? response.message : ''}</h1>
    </main>
  );
}

