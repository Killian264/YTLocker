import React, { useEffect } from "react";
import { LogoBar } from "../components/LogoBar";
import { RouteComponentProps, useLocation } from "react-router-dom";
import { useUserRefreshSession } from "../hooks/api/useUserRefreshSession";
import { GoogleOAuthCard } from "../components/GoogleOauthCard";

export interface LoginPageProps extends RouteComponentProps {
	className?: string;
}

export const LoginPage: React.FC<LoginPageProps> = ({ className, history }) => {
	const postSessionRefresh = useUserRefreshSession()

	let params = (new URLSearchParams(useLocation().search))
	let bearer = params.get("bearer")

	useEffect(() => {
		if(bearer != null){
			params.delete('bearer')
			history.replace({
			  search: params.toString(),
			})
	
			postSessionRefresh(bearer)
		}
	}, )

	return (
		<>
			<LogoBar className="absolute top-1 left-5"></LogoBar>
			<div className="flex h-screen">
				<div className="m-auto">
					<GoogleOAuthCard
						type="login"
					/>
				</div>
			</div>
		</>
	);
};

LoginPage.propTypes = {};
