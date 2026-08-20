import { MessagePage } from '../components/MessagePage';
import { messageMock } from '../lib/mock/store-and-serve-message';

export default function Home() {
  return <MessagePage response={messageMock} />;
}
