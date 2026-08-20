import { CenteredMessage } from "../components/CenteredMessage";
import { messageResponse } from "../lib/mock/render-centered-message";

export default function Home() {
  return <CenteredMessage response={messageResponse} />;
}
