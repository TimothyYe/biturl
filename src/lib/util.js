import moment from "moment";

export function timestampToString(timestamp) {
  return moment.unix(timestamp).format("L, LTS");
}
