"use client";

import { useEffect, useState } from "react";

import type { MessageResponse } from "../lib/mock/render-centered-message";
import styles from "./CenteredMessage.module.css";

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "/api";

export function CenteredMessage() {
  const [response, setResponse] = useState<MessageResponse>({ state: "loading" });

  useEffect(() => {
    const controller = new AbortController();

    fetch(`${apiBase}/v1/message`, { signal: controller.signal })
      .then(async (res) => {
        if (!res.ok) {
          throw new Error("request failed");
        }
        return res.json() as Promise<{ state?: string; message?: string }>;
      })
      .then((data) => {
        if (typeof data.message === "string" && data.message.trim()) {
          setResponse({ state: "success", message: data.message });
          return;
        }
        setResponse({ state: "empty" });
      })
      .catch(() => setResponse({ state: "error", error: { code: "INTERNAL", message: "request failed" } }));

    return () => controller.abort();
  }, []);

  if (response.state !== "success" || !response.message.trim()) {
    return <main className={styles.shell} aria-live="polite" />;
  }

  return (
    <main className={styles.shell} aria-live="polite">
      <h1 className={styles.message}>{response.message}</h1>
    </main>
  );
}
