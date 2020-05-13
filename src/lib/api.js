export function getRecords(provider, callback) {
  var url = "";
  if (provider) {
    url = "https://vpsdalao.com/api/records?provider=" + provider;
  } else {
    url = "https://vpsdalao.com/api/records";
  }

  fetch(url, {
    method: "GET"
  })
    .then(res => res.json())
    .then(content => callback(content))
    .catch(console.log);
}

export function getLastUpdate(callback) {
  fetch("https://vpsdalao.com/api/lastupdate", {
    method: "GET"
  })
    .then(res => res.json())
    .then(content => callback(content))
    .catch(console.log);
}

export default { getRecords, getLastUpdate };
