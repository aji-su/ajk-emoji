import { API_ENDPOINT } from "./constants.js";

export default async function(requestId) {
  const res = await fetch(`${API_ENDPOINT}/show/${requestId}`);
  const { emojis, originalImageUrl } = await res.json();
  return { emojis, originalImageUrl };
}
