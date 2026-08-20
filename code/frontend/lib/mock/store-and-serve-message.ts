export type MessageResponse =
  | { state: 'loading' }
  | { state: 'empty' }
  | { state: 'error'; error: { code: 'NOT_FOUND' | 'UNAVAILABLE' | 'INTERNAL'; message: string } }
  | { state: 'ready'; message: string };

export const messageMock: MessageResponse = {
  state: 'ready',
  message: 'Hello Word',
};
