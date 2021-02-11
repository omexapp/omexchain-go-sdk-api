package omex

/*
 OMEX general api
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

/*
 Time of the server running OMEX's REST API.
*/
func (client *Client) GetServerTime() (ServerTime, error) {
	var serverTime ServerTime
	_, err := client.Request(GET, OMEX_TIME_URI, nil, &serverTime)
	return serverTime, err
}
