import React from "react";

export interface CardProps {
	className?: string;
	children: React.ReactNode[];
}

export const Card: React.FC<CardProps> = ({ className, children }) => {
	let cardSize = "md:p-8 sm:p-6 p-4";

	return (
		<div className={`${className} ${cardSize} bg-primary-700 rounded-md`}>
			{children[0]}
			<div className="border-b-2 mt-3 mb-3"></div>
			{children.slice(1, children.length)}
		</div>
	);
};
