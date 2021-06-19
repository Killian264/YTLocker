import React, { useState } from "react";
import { Channel, Color } from "../shared/types";
import { ColorBadge } from "./ColorBadge";
import { Modal } from "./Modal";
import { RightArrow, Trash } from "./Svg";

export interface ChannelListItemProps {
	className?: string;
	channel: Channel;
	url: string;
	colors: Color[];
	mode: "normal" | "delete";
	remove: (channelId: string) => void;
}

const css = "rounded-md flex justify-between cursor-pointer";
const imageSize = "md:h-20 sm:h-16 h-16";
const textSize = "sm:text-md text-md";

export const ChannelListItem: React.FC<ChannelListItemProps> = ({
	channel,
	className,
	url,
	colors,
	mode,
	remove,
}) => {
	let [isOpen, setIsOpen] = useState(false);

	const hover = mode === "normal" ? "hover:bg-primary-600" : "hover:bg-red-700";

	const handleClick = () => {
		if (mode === "normal") {
			window.open(url, "_blank");
		}
		if (mode === "delete") {
			setIsOpen(true);
		}
	};

	const deleteClick = () => {
		remove(channel.youtubeId);
		setIsOpen(false);
	};

	const badges = colors.map((color, index) => {
		return <ColorBadge key={index} className="mt-2 mr-1" color={color}></ColorBadge>;
	});

	let trimmed =
		channel.description.length > 54 ? channel.description.substring(0, 54) + "..." : channel.description;

	return (
		<>
			{isOpen && (
				<Modal
					header={"Are you sure?"}
					body={"The channel will be removed form the playlist and new videos not added."}
					AcceptClick={deleteClick}
					RejectClick={() => {
						setIsOpen(false);
					}}
				/>
			)}
			<div className={`${className} ${css} ${hover}`} onClick={handleClick}>
				<div className="flex p-1">
					<img
						src={channel.thumbnailUrl}
						alt="Thumbnail"
						className={`rounded-lg object-cover ${imageSize}`}
					/>
					<div className="pl-3 flex flex-col">
						<span className={`${textSize} font-semibold`}>{channel.title}</span>
						<span>{trimmed}</span>
						<div>{badges}</div>
					</div>
				</div>
				<div className="mr-2 my-auto select-none">
					{mode === "normal" && <RightArrow size={24}></RightArrow>}
					{mode === "delete" && <Trash size={24}></Trash>}
				</div>
			</div>
		</>
	);
};
