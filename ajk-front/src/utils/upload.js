import { API_ENDPOINT } from "./constants.js";

export default async function(prefix, xsplit, imageAsDataUrl) {
  const body = JSON.stringify({
    prefix,
    xsplit,
    imageAsDataUrl
  });
  const res = await fetch(`${API_ENDPOINT}/create`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body
  });
  const { requestId } = await res.json();
  return requestId;
}
