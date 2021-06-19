import axios from "axios";
import { useEffect, useState } from "react";
import { BrowserRouter as Router, Redirect, Route, RouteProps, Switch, useHistory } from "react-router-dom";
import { DROPLET_BASE } from "../shared/env";
import { useBearer } from "../hooks/useBearer";
import { DashboardPage } from "./DashboardPage";
import { LoginPage } from "./LoginPage";
import { AlertProps } from "../components/Alert";
import { AlertContext } from "../hooks/AlertContext";
import { Alert } from "../components/Alert";

axios.defaults.baseURL = DROPLET_BASE;

const AuthenticatedRoute: React.FC<RouteProps> = ({ ...props }) => {
	const [loading, setLoading] = useState(true);
	const [bearer] = useBearer("");
	const history = useHistory();

	useEffect(() => {
		axios.defaults.headers.common["Authorization"] = bearer;

		let id = axios.interceptors.response.use(
			(response) => response,
			(error) => {
				if (typeof error.response != "undefined" && error.response.status === 401) {
					history.push("/login");
				}
				return error;
			}
		);

		setLoading(false);

		return () => {
			axios.interceptors.request.eject(id);
		};
	}, [bearer, history]);

	if (bearer === "") {
		return <Redirect to="/login" />;
	}

	if (loading) {
		return <div></div>;
	}

	return <Route {...props}></Route>;
};

const AlertsDisplay: React.FC<{}> = ({ children }) => {
	const [alert, _setAlert] = useState<AlertProps | null>(null);

	const setAlert = (alert: AlertProps) => {
		_setAlert(alert);
		setTimeout(() => {
			_setAlert(null);
		}, 4000);
	};

	return (
		<AlertContext.Provider value={{ pushAlert: setAlert }}>
			<div className="fixed w-full">{alert !== null && <Alert {...alert} />}</div>
			{children}
		</AlertContext.Provider>
	);
};

export const IndexPage: React.FC<{}> = () => {
	return (
		<AlertsDisplay>
			<Router>
				<Switch>
					<AuthenticatedRoute path="/" exact component={DashboardPage}></AuthenticatedRoute>
					<Route path="/login" exact component={LoginPage}></Route>
				</Switch>
			</Router>
		</AlertsDisplay>
	);
};
