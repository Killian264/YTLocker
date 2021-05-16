import { useState } from "react";
import { Playlist } from "../shared/types";
import { Card } from "./Card";
import { Modal } from "./Modal";
import { Cog, LeftArrow, SvgBox, Trash } from "./Svg";

export interface PlaylistViewProps {
	className?: string;
	playlist: Playlist;
	DeleteClick: () => void;
	BackClick: () => void;
}

export const PlaylistView: React.FC<PlaylistViewProps> = ({
	className,
	DeleteClick,
	BackClick,
	playlist,
}) => {
	let [isOpen, setIsOpen] = useState(false);

	const imageSize = "md:h-44 sm:h-32 h-24";

	const edit = () => {};

	const remove = () => {
		setIsOpen(true);
	};

	return (
		<>
			{isOpen && (
				<Modal
					header={"Are you sure?"}
					body={
						"Playlist information will be deleted, but the playlist will remain accessible on youtube."
					}
					AcceptClick={DeleteClick}
					RejectClick={() => {
						setIsOpen(false);
					}}
				/>
			)}
			<Card className={className}>
				<div className="flex justify-between -mb-1 -mt-1 items-center pb-1">
					<div className="flex items-center gap-2">
						<div onClick={BackClick} className="cursor-pointer -m-1">
							<LeftArrow size={32} strokeWidth={2}></LeftArrow>
						</div>
						<span className="leading-none text-2xl font-semibold">{playlist.title}</span>
					</div>
					<span className="text-sm leading-none pt-1">{`Created ${playlist.created.toDateString()}`}</span>
				</div>
				<div className="flex gap-2">
					<img
						src={playlist.thumbnail}
						alt="Logo"
						className={`col-span-6 rounded-lg object-cover ${imageSize}`}
					/>
					{/* TODO: Implement properly with grid */}
					<div className="flex-grow">
						<div className="md:flex-row flex-col flex gap-2 justify-between">
							<span className="md:text-3xl text-2xl font-semibold">{playlist.title}</span>
							<div className="flex gap-2">
								<SvgBox className="text-primary-200 border-primary-200" onClick={edit}>
									<Cog size={28}></Cog>
								</SvgBox>
								<SvgBox className="text-red-500 border-red-500" onClick={remove}>
									<Trash className="text-red-500" size={28}></Trash>
								</SvgBox>
							</div>
						</div>
						<div className="md:block hidden mt-2 ml-1">{playlist.description}</div>
					</div>
				</div>
				<div className="md:hidden sm:block mt-2 ml-1">{playlist.description}</div>
			</Card>
		</>
	);
};
