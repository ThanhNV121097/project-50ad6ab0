"use client";

import { useEffect, useState } from 'react';
import styles from './MessagePage.module.css';

type MessageResponse =
  | { state: 'loading' }
  | { state: 'empty' }
  | { state: 'error'; error: { code: 'NOT_FOUND' | 'UNAVAILABLE' | 'INTERNAL'; message: string } }
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
          if (res.status === 404) {
            setResponse({ state: 'empty' });
            return;
          }

          setResponse({
            state: 'error',
            error: {
              code: data?.error?.code ?? 'INTERNAL',
              message: data?.error?.message ?? 'Request failed',
            },
          });
          return;
        }

        setResponse({ state: 'ready', message: data.message });
      } catch {
        setResponse({ state: 'error', error: { code: 'INTERNAL', message: 'Request failed' } });
      }
    }

    loadMessage();
    return () => controller.abort();
  }, []);

  if (response.state === 'loading') return <main className={styles.shell} aria-busy="true" aria-live="polite"><p className={styles.message}>Loading</p></main>;
  if (response.state === 'error') return <main className={styles.shell} aria-live="polite"><p className={styles.message}>Error</p></main>;
  if (response.state === 'empty') return <main className={styles.shell} aria-live="polite"><p className={styles.message}>No message</p></main>;
  return <main className={styles.shell}><h1 className={styles.message}>{response.message}</h1></main>;
}
