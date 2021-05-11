import React from "react";
import { Channel } from "../shared/types";
import { Link } from "./Link";

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
			className={`${className} font-semibold text-xl hover:bg-secondary-hover rounded-md flex justify-between`}
		>
			<div className="flex">
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
			<div className="mr-3 text-3xl my-auto select-none">{">"}</div>
		</div>
	);
};
