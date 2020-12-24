exports.handler = async (event) => {
    console.info("EVENT\n" + JSON.stringify(event, null, 2))
  console.log("query string params ", event.queryStringParameters["hub.challenge"])
  return {
      'hub_challenge': event.queryStringParameters["hub.challenge"],
  };
};
