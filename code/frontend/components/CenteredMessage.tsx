"use client";

import type { MessageResponse } from "../lib/mock/render-centered-message";
import styles from "./CenteredMessage.module.css";

export function CenteredMessage({ response }: { response: MessageResponse }) {
  if (response.state !== "success" || !response.message.trim()) {
    return <main className={styles.shell} aria-live="polite" />;
  }

  return (
    <main className={styles.shell} aria-live="polite">
      <h1 className={styles.message}>{response.message}</h1>
    </main>
  );
}
