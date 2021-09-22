import React from "react";
import { Logo } from "./Logo";

export interface LogoBarProps {
	className?: string;
}

export const LogoBar: React.FC<LogoBarProps> = ({ className, children }) => {
	let pos = "inset-x-0 mx-auto mt-3";
	let css = `${className} ${pos} flex`;

	return (
		<div className={css}>
			<Logo></Logo>
			<span className="text-xl font-bold ml-1 text-accent tracking-wide">YTLocker</span>
		</div>
	);
};
