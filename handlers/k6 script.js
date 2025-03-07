import http from "k6/http";
import { check } from "k6";

// List of usernames and passwords
const users = {
  ajay: "ajaypwd",
  raghu: "raghupwd",
  susan: "susanpwd",
  mike: "mikepwd",
  john: "johnpwd",
  emily: "emilypwd",
  robert: "robertpwd",
  lisa: "lisapwd",
  peter: "peterpwd",
  anna: "annapwd",
  tom: "tompwd",
  julia: "juliapwd",
  kevin: "kevinpwd",
  sophie: "sophiepwd",
  david: "davidpwd",
  rachel: "rachelpwd",
  charles: "charlespwd",
  mary: "marypwd",
  paul: "paulpwd",
  steve: "stevepwd",
  jack: "jackpwd",
  james: "jamespwd",
  karen: "karenpwd",
  sandra: "sandrapwd",
  ben: "benpwd",
  olivia: "oliviapwd",
  matt: "mattpwd",
  laura: "laurapwd",
  will: "willpwd",
  claire: "clairepwd",
  alex: "alexpwd",
  grace: "gracepwd",
  ian: "ianpwd",
  zoe: "zoepwd",
  daniel: "danielpwd",
  lucy: "lucypwd",
  george: "georgepwd",
  beth: "bethpwd",
  tim: "timpwd",
  isla: "islapwd",
  jim: "jimpwd",
  sarah: "sarahpwd",
  kyle: "kylepwd",
  vicky: "vickypwd",
  maggie: "maggiepwd",
  leah: "leahpwd",
  max: "maxpwd",
  joseph: "josephpwd",
  luke: "lukepwd",
  holly: "hollypwd",
  carl: "carlpwd",
  donna: "donnapwd",
  jenny: "jennypwd",
  greg: "gregpwd",
  amy: "amypwd",
  fiona: "fionapwd",
  brian: "brianpwd",
  toni: "tonipwd",
  lucas: "lucaspwd",
  olga: "olgapwd",
  morgan: "morganpwd",
  mia: "miapwd",
  nathan: "nathanpwd",
  ella: "ellapwd",
  sam: "sampwd",
  geoff: "geoffpwd",
  tina: "tinapwd",
  victor: "victorpwd",
  leon: "leonpwd",
  alison: "alisonpwd",
  hannah: "hannahpwd",
  roxy: "roxypwd",
  chris: "chrispwd",
  deborah: "deborahpwd",
  mark: "markpwd",
  nina: "ninapwd",
  hugo: "hugopwd",
  tess: "tesspwd",
  daisy: "daisypwd",
  aiden: "aidenpwd",
  rose: "rosepwd",
  simon: "simonpwd",
  helen: "helenpwd",
  natalie: "nataliepwd",
  chloe: "chloepwd",
  jackson: "jacksonpwd",
  quinn: "quinnpwd",
  penny: "pennypwd",
  ellen: "ellenpwd",
  sean: "seanpwd",
  toby: "tobypwd",
  jason: "jasonpwd",
  felix: "felixpwd",
  kim: "kimpwd",
  maria: "mariapwd",
  jasmine: "jasminepwd",
  josh: "joshpwd",
  jennifer: "jenniferpwd",
  ron: "ronpwd",
  alina: "alinapwd",
  edward: "edwardpwd",
  jacob: "jacobpwd",
  valerie: "valeriepwd",
  hank: "hankpwd",
  nancy: "nancypwd",
};

export let options = {
  stages: [
    { duration: "5m", target: 100 }, // Ramp up to 100 users over 5 minutes
  ],
};

export default function () {
  // Randomly pick a user
  const userKeys = Object.keys(users);
  const randomIndex = Math.floor(Math.random() * userKeys.length);
  const user = userKeys[randomIndex];
  const password = users[user];

  // The API URL for token creation
  const tokenApiUrl = "http://localhost:8080/login";

  // Prepare the payload for authentication (adjust as necessary for your API)
  const payload = JSON.stringify({
    username: user,
    password: password,
  });

  // Set the headers for the request
  const headers = {
    "Content-Type": "application/json",
  };

  // Send a POST request to the token creation API
  let res = http.post(tokenApiUrl, payload, { headers });

  // Check if the request was successful (status 200)
  check(res, {
    "Token created successfully": (r) => r.status === 200,
  });

  // Optionally log the response or add further checks if needed
  // console.log('Response: ' + JSON.stringify(res.json()));
}
