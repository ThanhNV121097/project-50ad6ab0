"use client";

import { useEffect, useState } from 'react';
import styles from './MessagePage.module.css';

type MessageResponse = { message: string };

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? '/api';

export function MessagePage() {
  const [message, setMessage] = useState<string | null>(null);

  useEffect(() => {
    const controller = new AbortController();

    async function loadMessage() {
      try {
        const res = await fetch(`${apiBase}/v1/message`, {
          signal: controller.signal,
          headers: { Accept: 'application/json' },
        });

        const data = (await res.json().catch(() => null)) as MessageResponse | null;

        if (!res.ok || typeof data?.message !== 'string') {
          setMessage(null);
          return;
        }

        setMessage(data.message);
      } catch {
        setMessage(null);
      }
    }

    loadMessage();

    return () => controller.abort();
  }, []);

  return (
    <main className={styles.shell}>
      {message ? <h1 className={styles.message}>{message}</h1> : null}
    </main>
  );
}

