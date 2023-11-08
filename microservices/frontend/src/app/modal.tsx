const Modal = ({ isOpen, onClose, children }) => {
    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center overflow-x-hidden overflow-y-auto outline-none">
            <div className="modal-overlay" onClick={onClose} />
            <div className="modal-container bg-white w-96 md:w-1/2 mx-auto rounded shadow-lg z-50">
                <div className="modal-content py-4 text-left px-6">{children}</div>
            </div>
        </div>
    );
};

export default Modal;
