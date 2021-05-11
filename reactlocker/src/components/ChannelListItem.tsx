import React from "react";
import { Channel } from "../shared/types";
import { Link } from "./Link";
import { Arrow } from "./Svg";

export interface ChannelListItemProps {
	className?: string;
	channel: Channel;
}

export const ChannelListItem: React.FC<ChannelListItemProps> = ({
	channel,
	className,
}) => {
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between`;

	const imageSize = "md:h-20 sm:h-16 h-16";

	const textSize = "sm:text-lg text-md";

	return (
		<div className={css}>
			<div className="flex p-1">
				<img
					src={channel.thumbnail}
					alt="Logo"
					className={`rounded-lg object-cover ${imageSize}`}
				/>
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold`}>
						{channel.title}
					</span>
					<Link
						className={`${textSize} text-accent ml-0.5`}
						href={channel.url}
						target="_blank"
					>
						Youtube Link
					</Link>
				</div>
			</div>
			<div className="mr-2 my-auto select-none">
				<Arrow size={24}></Arrow>
			</div>
		</div>
	);
};