import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import { DashboardPage } from "./DashboardPage";
import { LoginPage } from "./LoginPage";

export const IndexPage: React.FC<{}> = () => {
	return (
		<Router>
			<Switch>
				<Route path="/" exact component={DashboardPage}></Route>
				<Route path="/login" exact component={LoginPage}></Route>
			</Switch>
		</Router>
	);
};
