import React from "react";
import logo from "../static/logo.png";

export interface LogoProps {
	className?: string;
}

export const Logo: React.FC<LogoProps> = () => {
	return <img src={logo} alt="Logo" width="32" height="32" />;
};
