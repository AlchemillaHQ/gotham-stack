import Swal from 'sweetalert2/src/sweetalert2.js'
import DOMPurify from 'isomorphic-dompurify';
import '../../css/sweetalert2.scss';

window.Swal = Swal;
window.htmx = require("htmx.org");
window.DOMPurify = DOMPurify;