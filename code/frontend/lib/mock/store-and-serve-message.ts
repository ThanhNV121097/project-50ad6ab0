export type MessageResponse =
  | { state: 'loading' }
  | { state: 'error'; error: { code: 'INTERNAL' | 'UNAVAILABLE' | 'NOT_FOUND'; message: string } }
  | { state: 'empty' }
  | { state: 'ready'; message: string };

export const messageMock: MessageResponse = {
  state: 'ready',
  message: 'Hello Word',
};
