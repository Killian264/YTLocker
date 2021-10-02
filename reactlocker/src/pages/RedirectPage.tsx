import React, { useContext, useEffect } from "react";
import { LogoBar } from "../components/LogoBar";
import { RouteComponentProps, useLocation } from "react-router-dom";
import { useUserRefreshSession } from "../hooks/api/useUserRefreshSession";
import { AlertContext } from "../hooks/AlertContext";

export interface RedirectPageProps extends RouteComponentProps {
	className?: string;
}

export const RedirectPage: React.FC<RedirectPageProps> = ({ className, history }) => {
	const { pushAlert } = useContext(AlertContext);
	const refreshSession = useUserRefreshSession()

	let params = (new URLSearchParams(useLocation().search))
	let success = params.get("success") === "true"
	let reason = params.get("reason")
	let bearer = params.get("bearer")

	useEffect(() => {
		if(bearer !== null){
			refreshSession(bearer)
			return;
		}

		if(!success && reason !== null){
			pushAlert({
				message: reason,
				type: "failure",
			})
		}

		if(success){
			pushAlert({
				message: "Successfully linked account.",
				type: "success",
			})
		}

		history.push("/")
	}, [bearer, reason, success, pushAlert, history, refreshSession])

	return (
		<>
			<LogoBar className="absolute top-1 left-5"></LogoBar>
			<div className="flex h-screen">
				
			</div>
		</>
	);
};