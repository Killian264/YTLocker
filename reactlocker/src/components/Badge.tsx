import React from "react";

export interface BadgeProps {
	className?: string;
	color: keyof typeof colors;
}

const colors = {
	primary: "text-accent-text bg-accent disabled:bg-accent-disabled disabled:text-accent-text-disabled",
};

export const Badge: React.FC<BadgeProps> = ({ className, color, children }) => {
	let size = "px-1 text-xs rounded";

	return <span className={`${className} ${size} ${colors[color]} font-bold select-none`}>{children}</span>;
};
