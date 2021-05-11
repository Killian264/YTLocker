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
	return (
		<div
			className={`${className} font-semibold text-xl hover:bg-primary-600 rounded-md flex justify-between`}
		>
			<div className="flex p-1">
				<img
					src={channel.thumbnail}
					alt="Logo"
					width="140"
					height="120"
					className="rounded-lg"
				/>
				<div className="pl-3 flex flex-col">
					<span>{channel.title}</span>
					<Link
						className="text-accent text-lg ml-0.5"
						href={channel.url}
						target="_blank"
					>
						Youtube Link
					</Link>
				</div>
			</div>
			<div className="mr-3 my-auto select-none">
				<Arrow size={24}></Arrow>
			</div>
		</div>
	);
};
