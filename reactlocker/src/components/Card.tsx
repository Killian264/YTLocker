import React from "react";

export interface CardProps {
	className?: string;
	children: React.ReactNode[];
}

export const Card: React.FC<CardProps> = ({ className, children }) => {
	return (
		<div className={`${className} bg-secondary p-8 rounded-md`}>
			{children[0]}
			<div className="border-b-2 mt-3 mb-3"></div>
			{children.slice(1, children.length)}
		</div>
	);
};
