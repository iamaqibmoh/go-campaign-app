fetch("http://localhost:2802/api/v1/campaigns/5")
    .then(res => res.json())
    .then(data => console.log(data))
