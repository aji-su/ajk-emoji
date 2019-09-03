import { API_ENDPOINT } from "./constants.js";

export default async function(requestId) {
  const res = await fetch(`${API_ENDPOINT}/download/${requestId}`);
  const { url } = await res.json();
  return url;
}
