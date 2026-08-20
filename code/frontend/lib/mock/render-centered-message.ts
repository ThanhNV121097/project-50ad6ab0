export type MessageResponse =
  | { state: 'loading' }
  | { state: 'error'; error: { code: 'INTERNAL'; message: string } }
  | { state: 'empty' }
  | { state: 'success'; message: string };

export const messageResponse: MessageResponse = {
  state: 'success',
  message: 'Hello Word',
};
