import React from "react";
import { ColorBadge } from "./ColorBadge";
import { RightArrow } from "./Svg";

export interface LoadingListItemProps {
	className?: string;
}

export const LoadingListItem: React.FC<LoadingListItemProps> = ({ className }) => {
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between cursor-pointer`;
	const imageSize = "md:h-20 md:w-32 h-16 w-24";
	const textSize = "sm:text-md text-md";

	const color = "bg-primary-600 text-primary-600 rounded leading-5";

	return (
		<div className={css}>
			<div className="flex p-1 animate-pulse overflow-hidden">
				<div className={`${color} ${imageSize} flex-shrink-0 rounded-lg object-cover`} />
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} ${color} font-semibold whitespace-nowrap`}>
						sadfasdasdfasdfsdfsdfasdf
					</span>
					<div>
						<span className={`${color} text-sm mt-1`}>asddsfasdfasdffasdf</span>
					</div>
					<div>
						<ColorBadge className={color} color="blue-1"></ColorBadge>
					</div>
				</div>
			</div>
			<div className="mr-2 my-auto select-none">
				<RightArrow size={24}></RightArrow>
			</div>
		</div>
	);
};
